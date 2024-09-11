import { ListingSearchInput } from "../../zod/paginate";
import api from "../axios";
import { narrowErrorMessage } from "../error";
import { GenericRes } from "../types";

export type ListingRecord = {
  id: string;
  listingUrl: string;
  name: string;
  summary: string;
  images: string;
  address: string;
  bedrooms: number;
  beds: number;
  roomType: string;
  bedType: string;
  price: string;
};

type CategoryType = {
  roomType: Array<string>;
  bedType: Array<string>;
  propertyType: Array<string>;
};

export const listingRecordsApi = async (query: ListingSearchInput) => {
  try {
    let params = {};

    if (query) {
      params = {
        ...query,
      };
    }
    return await api.get<GenericRes<ListingRecord[]>>("/listings", { params });
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};

export const categoryTypeApi = async () => {
  try {
    return await api.get<GenericRes<CategoryType>>("/listings/categories");
  } catch (error) {
    throw new Error(narrowErrorMessage(error));
  }
};
