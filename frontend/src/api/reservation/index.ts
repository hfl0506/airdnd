import { ReservationInput } from "../../zod/reservation";
import api from "../axios";
import { narrowErrorMessage } from "../error";
import { GenericRes } from "../types";

export const createReservationApi = async (payload: ReservationInput) => {
  try {
    return await api.post<GenericRes<string>>("/bookings", payload);
  } catch (error) {
    throw new Error(narrowErrorMessage(error))
  }
}

