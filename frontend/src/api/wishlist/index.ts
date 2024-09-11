import api from "../axios";
import { narrowErrorMessage } from "../error";
import { GenericRes } from "../types";
import { RoomRes } from "../rooms";
import { UpsertWishlistInput } from "../../zod/wishlist";

type WishlistRes = {
  id: string;
  room: RoomRes;
  createdAt: Date;
  updatedAt: Date;
};

export const getWishlistApi = async () => {
  try {
    return await api.get<GenericRes<WishlistRes[]>>("/wishlist");
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const getWishlistIdsApi = async () => {
  try {
    return await api.get<GenericRes<string[]>>("/wishlist/ids");
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const upsertWishlist = async (payload: UpsertWishlistInput) => {
  try {
    return await api.post<GenericRes<string>>("/wishlist", payload)
  } catch (error) {
    throw new Error(narrowErrorMessage(error))
  }
}
