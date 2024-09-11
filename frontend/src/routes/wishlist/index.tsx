import { createFileRoute } from "@tanstack/react-router";
import { getWishlistApi, upsertWishlist } from "../../api/wishlist";
import { Delete } from "lucide-react";
import { useMutation } from "@tanstack/react-query";
import { UpsertWishlistInput } from "../../zod/wishlist";

export const Route = createFileRoute("/wishlist/")({
  loader: () => getWishlistApi(),
  component: WishList,
});

function WishList() {
  const wishlist = Route.useLoaderData();
  const navigate = Route.useNavigate();
  const listingMutation = useMutation({
    mutationKey: ["wishlist_delete"],
    mutationFn: (payload: UpsertWishlistInput) => upsertWishlist(payload),
    onSuccess: () => {
      navigate({ to: "/wishlist" })
    }
  })

  const handleDelete = (roomId: string) => {
    listingMutation.mutate({ roomId, liking: false });
  }

  return <div className="w-full min-h-screen px-10 mt-10 flex flex-col gap-4">
    <h1>Room Wishlist: </h1>
    <table className="border-slate-100 border-[1px] rounded-md p-4 text-center">
      <tr className="border-b-[1px] border-slate-100">
        <th>Name</th>
        <th>Room Type</th>
        <th>Bed Type</th>
        <th>Bed Rooms</th>
        <th>Beds</th>
        <th>Bath Rooms</th>
        <th>Price</th>
        <th>Action</th>
      </tr>
      {wishlist.data.data?.length > 0 ? wishlist.data.data.map(item =>
        <tr className="text-md font-medium">
          <td>{item.room.name}</td>
          <td>{item.room.roomType}</td>
          <td>{item.room.bedType}</td>
          <td>{item.room.bedrooms}</td>
          <td>{item.room.beds}</td>
          <td>{item.room.bathrooms}</td>
          <td>{item.room.price}</td>
          <td><Delete onClick={(() => handleDelete(item.room.id))} className="mx-auto cursor-pointer" /></td>
        </tr>
      ) : <tr className="w-full text-center"><td>N/A</td></tr>}
    </table>
  </div>;
}
