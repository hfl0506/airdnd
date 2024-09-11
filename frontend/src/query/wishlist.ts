import { queryOptions } from "@tanstack/react-query";
import { getWishlistIdsApi } from "../api/wishlist";

export const wishlistIdsQueryOptions = () => {
  return queryOptions({
    queryKey: ["wishlist_liking"],
    queryFn: () => getWishlistIdsApi()
  })
}
