import api from "../axios";
import { narrowErrorMessage } from "../error";
import { RoomRes } from "../rooms";
import { GenericRes } from "../types";

export type Booking = {
  id: string;
  room: RoomRes;
  checkIn: {
    adult: number;
    child: number;
    infant: number;
  },
  price: number;
  startDate: Date;
  endDate: Date;
  createdAt: Date;
  updatedAt: Date;
}

export const listBookingsApi = async () => {
  try {
    return await api.get<GenericRes<Booking[]>>("/bookings");
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
}

export const deleteBookingApi = async (id: string) => {
  try {
    return await api.delete<GenericRes<string>>(`/bookings/${id}`);
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
}
