import { LoginInput, RegisterInput } from "../../zod/auth";
import api from "../axios";
import { GenericRes } from "../types";
import { narrowErrorMessage } from "../error";

export type User = {
  name: string;
  email: string;
  photo: string | null;
};

type AuthRes = {
  sessionId: string;
  accessToken: string;
  refreshToken: string;
  accessTokenExpiresAt: Date;
  refreshTokenExpiresAt: Date;
  user: User;
};

type RefreshRes = {
  accessToken: string;
  accessTokenExpiresAt: Date;
};

export const loginApi = async (payload: LoginInput) => {
  try {
    return await api.post<GenericRes<AuthRes>>("/users/login", payload);
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const refreshApi = async (refreshToken: string) => {
  try {
    return await api.post<GenericRes<RefreshRes>>("/users/refresh", {
      refreshToken,
    });
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const registerApi = async (payload: RegisterInput) => {
  try {
    return await api.post<GenericRes<AuthRes>>("/users/register", payload);
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const getMeApi = async () => {
  try {
    return await api.get<GenericRes<User>>("/users/me");
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};
