const ACCESS_TOKEN = "access_token";
const REFRESH_TOKEN = "refresh_token";

type Tokens = "access_token" | "refresh_token";

export function setToken(key: Tokens, value: string) {
  localStorage.setItem(key, value);
}

export function setTokens(at: string, rt: string) {
  setToken(ACCESS_TOKEN, at);
  setToken(REFRESH_TOKEN, rt);
}

export function removeTokens() {
  localStorage.removeItem(ACCESS_TOKEN);
  localStorage.removeItem(REFRESH_TOKEN);
}

export function getToken(key: Tokens) {
  return localStorage.getItem(key) ?? "";
}
