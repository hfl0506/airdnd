import { AxiosError } from "axios";

export function narrowErrorMessage(error: unknown) {
  if (error instanceof Error) return error.message;
  else if (error instanceof AxiosError) return error.message;
  else return `${error}`;
}
