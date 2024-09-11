import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { categoryQueryOptions, useInfiniteScroll } from "../query/listing";
import { useMutation, useSuspenseQuery } from "@tanstack/react-query";
import { listingSearchSchema } from "../zod/paginate";
import Card from "../components/card";
import Tabs from "../components/tabs";
import CustomImage from "../components/image";
import LikeIcon from "../components/like-icon";
import { UpsertWishlistInput } from "../zod/wishlist";
import { upsertWishlist } from "../api/wishlist";
import { wishlistIdsQueryOptions } from "../query/wishlist";
import { useInView } from "react-intersection-observer";
import { useEffect } from "react";

export const Route = createFileRoute("/")({
  validateSearch: listingSearchSchema,
  loader: async ({ context: { queryClient } }) => {
    return await queryClient.ensureQueryData(categoryQueryOptions());
  },
  component: Index,
  notFoundComponent: () => {
    return <div>Not found</div>;
  },
});

function Index() {
  const { ref, inView } = useInView();
  const { tab_id } = Route.useSearch();
  const navigate = useNavigate({ from: Route.fullPath });
  const { data: categories } = Route.useLoaderData();
  const { data: roomIds, refetch: refetchRoomIds } = useSuspenseQuery(wishlistIdsQueryOptions());
  const { data: listings, loading, errorMessage, hasNextPage, fetchNextPage, refetch: refetchListing } = useInfiniteScroll({
    tabId: tab_id
  })

  const likingMutation = useMutation({
    mutationKey: ["liking"],
    mutationFn: (payload: UpsertWishlistInput) => upsertWishlist(payload),
    onSuccess: async () => {
      await refetchRoomIds();
    }
  })

  const activeTabClick = (tab: string) => {
    navigate({
      search: (prev) => ({ ...prev, tab_id: tab }),
    });
    refetchListing();
  };



  const handleLiking = async (roomId: string) => {
    const hasRoomId = roomIds.data.data.includes(roomId)
    likingMutation.mutate({ roomId, liking: !hasRoomId })
  }

  useEffect(() => {
    if (inView && hasNextPage) {
      fetchNextPage();
    }
  }, [inView, hasNextPage])

  return (
    <div className="flex flex-col min-h-screen gap-10 mt-10">
      <Tabs
        tabs={categories.data.propertyType}
        activeTab={tab_id ?? ""}
        onClick={activeTabClick}
      />
      <input className="w-full h-12 max-w-[300px] mx-auto rounded-md border-slate-500 border-[1px] indent-2" />
      <div className="flex flex-wrap w-full max-w-[1600px] my-0 mx-auto gap-4 justify-center">
        {listings?.length > 0
          ? listings.map(room =>
            <Card
              key={room.id}
              onClick={() =>
                navigate({
                  to: "/rooms/$roomId",
                  params: { roomId: room.id },
                })
              }
            >
              <LikeIcon isIncluded={roomIds.data.data.includes(room.id)} onClick={(e) => {
                e.stopPropagation();
                handleLiking(room.id);
              }} />
              <CustomImage
                className="w-full h-[200px] rounded-lg object-fill"
                src={room.images}
              />
              <h1 className="font-bold">{room.address}</h1>
              <h2 className="text-slate-500">Beds: {room.beds}</h2>
              <h1 className="font-bold">
                {room.price} CAD <span>night</span>
              </h1>
            </Card>
          )
          : null}
        <div ref={ref}></div>
      </div>
    </div>
  );
}
