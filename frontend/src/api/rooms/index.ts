import { narrowErrorMessage } from "../error";
import api from "../axios";
import { GenericRes } from "../types";

export type ReviewDetail = {
  id: string;
  user: {
    name: string;
    email: string;
    photo: string | null;
  },
  comment: string;
  rating: number;
  createdAt: Date;
  updatedAt: Date;
}

export type RoomRes = {
  id: string;
  name: string;
  summary: string;
  space: string;
  description: string;
  propertyType: string;
  roomType: string;
  bedType: string;
  accommodates: number;
  bedrooms: number;
  beds: number;
  bathrooms: string;
  amenities: string[];
  price: string;
  image: string;
  address: string;
  reviews: ReviewDetail[]
};

export const getRoomByIdApi = async (id: string) => {
  try {
    return await api.get<GenericRes<RoomRes>>(`/rooms/${id}`);
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};
