import { queryOptions } from "@tanstack/react-query";
import { listBookingsApi } from "../api/bookings";

export const bookingQueryOptions = () => {
  return queryOptions({
    queryKey: ['bookings'],
    queryFn: () => listBookingsApi()
  })
}
