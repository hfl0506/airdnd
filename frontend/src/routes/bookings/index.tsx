import { useMutation, useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { bookingQueryOptions } from "../../query/booking";
import { parseDatetimeFormat } from "../../utils/day";
import { deleteBookingApi } from "../../api/bookings";
import { DeleteIcon } from "lucide-react";

export const Route = createFileRoute("/bookings/")({
  component: Bookings,
});

function Bookings() {
  const { data: bookings, refetch } = useSuspenseQuery(bookingQueryOptions());
  const bookingMutation = useMutation({
    mutationKey: ["booking_delete"],
    mutationFn: (id: string) => deleteBookingApi(id),
    onSuccess: () => {
      refetch();
    }
  })

  const handleDelete = (id: string) => {
    bookingMutation.mutate(id);
  }
  return <div className="flex flex-col w-full min-h-screen items-center max-w-[1400px] mt-10 mx-auto">
    <table className="border-slate-200 border-[1px] w-full text-center">
      <tr>
        <th>Name</th>
        <th>Check in</th>
        <th>Start Date</th>
        <th>End Date</th>
        <th>Price</th>
        <th>Action</th>
      </tr>
      {bookings.data.data?.length > 0 ? bookings.data.data.map(booking => <tr>
        <td>{booking.room.name}</td>
        <td>Adult: {booking.checkIn.adult} | Child: {booking.checkIn.child} | Infant: {booking.checkIn.infant}</td>
        <td>{parseDatetimeFormat(booking.startDate)}</td>
        <td>{parseDatetimeFormat(booking.endDate)}</td>
        <td>{booking.price}</td>
        <td><DeleteIcon onClick={() => handleDelete(booking.id)} className="cursor-pointer m-auto" /></td>
      </tr>) : <tr><td>N/A</td></tr>}
    </table>
  </div>
}
