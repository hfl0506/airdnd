import { CreateReviewInput } from "../../zod/review"
import api from "../axios"
import { narrowErrorMessage } from "../error"
import { GenericRes } from "../types"

export const createReviewApi = async (payload: CreateReviewInput) => {
  try {
    return await api.post<GenericRes<string>>("/reviews", payload);
  } catch (error) {
    throw new Error(narrowErrorMessage(error))
  }
}
