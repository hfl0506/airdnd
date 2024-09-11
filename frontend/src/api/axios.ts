import axios, { AxiosError, InternalAxiosRequestConfig } from "axios";
import { getToken, removeTokens, setToken } from "../utils/token";
import { refreshApi } from "./auth";

interface FailedQueueItem {
  resolve: (value?: unknown) => void;
  reject: (reason?: unknown) => void;
}

interface CustomAxiosRequestConfig extends InternalAxiosRequestConfig {
  _retry?: boolean;
}

const BASE_URL = import.meta.env.VITE_BASE_API_URL;

const instance = axios.create({
  baseURL: BASE_URL,
  timeout: 5000,
});

let isRefreshing = false;
let failedQueue: FailedQueueItem[] = [];

const processQueue = (
  error: AxiosError | null,
  token: string | null = null
) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });

  failedQueue = [];
};

instance.interceptors.request.use(
  function (config: CustomAxiosRequestConfig) {
    const token = getToken("access_token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

instance.interceptors.response.use(
  function (response) {
    return response;
  },
  async function (error: AxiosError) {
    const originalRequest = error.config as CustomAxiosRequestConfig;
    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        try {
          const token = await new Promise((resolve, reject) => {
            failedQueue.push({ resolve, reject });
          });
          originalRequest.headers.Authorization = `Bearer ${token}`;
          return await instance(originalRequest);
        } catch (err) {
          return await Promise.reject(err);
        }
      }

      originalRequest._retry = true;
      isRefreshing = true;

      const refreshToken = getToken("refresh_token");

      return new Promise((resolve, reject) => {
        refreshApi(refreshToken)
          .then((res) => {
            const newAccessToken = res.data.data.accessToken;
            setToken("access_token", newAccessToken);
            instance.defaults.headers.common.Authorization = `Bearer ${newAccessToken}`;
            processQueue(null, newAccessToken);
            resolve(instance(originalRequest));
          })
          .catch((err) => {
            processQueue(err, null);
            removeTokens();
            reject(err);
          })
          .finally(() => {
            isRefreshing = false;
          });
      });
    }

    return Promise.reject(error);
  }
);

export default instance;
