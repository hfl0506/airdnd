import { createFileRoute } from "@tanstack/react-router";
import { getRoomByIdApi } from "../../../api/rooms";
import AmenityMap from "../../../components/amenityMap";
import CustomImage from "../../../components/image";
import { useEffect, useRef, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { ReservationInput, ReservationSchema, reservationSchema } from "../../../zod/reservation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { createReservationApi } from "../../../api/reservation";
import DatePicker from "react-datepicker";
import Typography from "../../../components/typography";
import Button from "../../../components/button";
import "react-datepicker/dist/react-datepicker.css"
import Label from "../../../components/label";
import { parseDatetimeFormat } from "../../../utils/day";
import Avatar from "../../../components/avatar";
import { Star } from "lucide-react";
import ReviewForm from "../../../components/reviewForm";
import { CreateReviewInput } from "../../../zod/review";
import { createReviewApi } from "../../../api/review";

const DAY_TIME = 1000 * 60 * 60 * 24;

export const Route = createFileRoute("/rooms/$roomId/")({
  loader: async ({ params }) => {
    return await getRoomByIdApi(params.roomId);
  },
  component: RoomId,
});

function RoomId() {
  const { data } = Route.useLoaderData();
  const navigate = Route.useNavigate();
  const [open, setOpen] = useState(false);
  const modalRef = useRef<HTMLDivElement>(null)
  const { register, watch, setValue, control, handleSubmit, reset, formState: { errors, isLoading } } = useForm<ReservationSchema>({
    defaultValues: {
      startDate: undefined,
      endDate: undefined,
      roomId: data.data.id ?? "",
      price: parseInt(data.data.price) ?? 0,
      adult: "0",
      child: "0",
      infant: "0"
    },
    resolver: zodResolver(reservationSchema)
  })
  const reservationMutation = useMutation({
    mutationFn: (data: ReservationInput) => createReservationApi(data),
    mutationKey: ["reservation"],
    onSuccess: () => {
      reset();
      setOpen(false);
    },
    onError: (error) => {
      console.error(error)
    }
  })
  const [startDate, endDate, price] = watch(["startDate", "endDate", "price"])
  const reviewMutation = useMutation({
    mutationFn: (data: CreateReviewInput) => createReviewApi(data),
    mutationKey: ["review"],
    onSuccess: () => {
      navigate({ to: `/rooms/$roomId`, params: { roomId: data.data.id } })
    }
  })

  const openReservationModal = () => {
    setOpen(true)
  }

  const handleOutsideClick = (e: MouseEvent | TouchEvent) => {
    if (modalRef.current && !modalRef.current.contains(e.target as Node)) {
      setOpen(false);
    }
  }

  const onSubmit = handleSubmit((data) => {
    reservationMutation.mutate({
      startDate: parseDatetimeFormat(data.startDate),
      endDate: parseDatetimeFormat(data.endDate),
      roomId: data.roomId,
      price: data.price,
      checkIn: {
        adult: parseInt(data.adult),
        child: parseInt(data.child),
        infant: parseInt(data.infant)
      }
    });
  })

  const reviewCallback = (data: CreateReviewInput) => {
    reviewMutation.mutate(data);
  }

  useEffect(() => {
    if (startDate && endDate) {
      const timeDiff = endDate.getTime() - startDate.getTime();
      const dayDiff = Math.floor(timeDiff / DAY_TIME);
      if (!isNaN(dayDiff)) {
        setValue("price", parseInt(data.data.price) * dayDiff);
      }
    }
  }, [startDate, endDate])


  useEffect(() => {
    document.addEventListener("mouseup", handleOutsideClick);
    document.addEventListener("touchend", handleOutsideClick)
    return () => {
      document.removeEventListener("mouseup", handleOutsideClick)
      document.removeEventListener("touchend", handleOutsideClick)
    }
  }, [])
  return (
    <div className="w-full min-h-screen flex flex-col gap-4 px-10 my-10 mx-auto md:max-w-[1600px]">
      <h1 className="text-2xl">{data.data.name}</h1>
      <CustomImage
        src={data.data.image}
        loading="lazy"
        className="w-full h-[250px] rounded-md md:h-[400px] xl:h-[500px] xl:w-[80%]"
      />
      <h2 className="text-lg">{data.data.address}</h2>
      <div className="flex items-center gap-2">
        <span>
          {data.data.accommodates}{" "}
          {data.data.accommodates > 1 ? "guests" : "guest"}
        </span>
        <span>-</span>
        <span>{data.data.bedrooms} bedrooms</span>
        <span>-</span>
        <span>{data.data.beds} beds</span>
        <span>-</span>
        <span>{parseInt(data.data.bathrooms)} baths</span>
      </div>
      <hr />
      <p>{data.data.description}</p>
      <hr />
      <div className="flex flex-col gap-2">
        <h1>What this place offers</h1>
        {data.data.amenities?.length > 0
          ? data.data.amenities.map((amenity) => (
            <AmenityMap amenity={amenity} />
          ))
          : "N/A"}
      </div>
      <button type='button' onClick={openReservationModal} className="w-full max-w-[200px] bottom-0 sticky h-[40px] bg-red-300 px-8 py-4 text-white rounded-md flex items-center justify-center">Reservation</button>
      <label>Reviews:</label>
      {
        data.data.reviews?.length > 0 ? <>
          {data.data.reviews.map(review => <div className="min-h-20 w-full max-w-[350px] flex flex-col justify-start gap-2 border-[1px] border-slate-200 rounded-md shadow-sm px-4 py-2">
            <h1 className="flex items-center gap-2">{review.user.name} | <Avatar src={review.user.photo!} name={review.user.name} /> | {parseDatetimeFormat(review.updatedAt)}</h1>
            <p className="text-slate-500 text-md">{review.comment}</p>
            <p className="flex items-center">{new Array(review.rating).fill(0).map(() => <Star className="text-yellow-500" />)}</p>
          </div>)}
        </> : null}
      <ReviewForm roomId={data.data.id} callback={reviewCallback} />
      {open &&
        <div
          ref={modalRef}
          className="absolute top-1/2 left-1/2 -translate-x-1/2 translate-y-1/2 w-full max-w-[450px] rounded-md min-h-[350px] shadow-md bg-white flex px-4 py-6 flex-col items-center gap-2"
        >
          <h1 className="text-xl">Reservation</h1>
          <form onSubmit={onSubmit} className="w-full flex flex-col h-full gap-2">

            <Controller control={control} name="startDate" render={({ field }) => (
              <Label>
                Start Date:
                <DatePicker
                  placeholderText="Select Start Date"
                  onChange={(date) => field.onChange(date)}
                  selected={field.value}
                />
                {errors.startDate?.message && <Typography>{errors.startDate.message}</Typography>}
              </Label>
            )}
            />
            <Controller
              control={control}
              name="endDate"
              render={({ field }) => (
                <Label>
                  End Date:
                  <DatePicker
                    placeholderText="Select End Date"
                    onChange={(date) => field.onChange(date)}
                    selected={field.value}
                  />
                  {errors.endDate?.message && <Typography>{errors.endDate?.message}</Typography>}
                </Label>

              )}
            />

            <Label>
              Price:
              <span>{price} CAD</span>
              {errors.price?.message && <Typography>{errors.price.message}</Typography>}
            </Label>
            <Label>
              Adults:
              <input type="number" min="0" {...register("adult")} />
              {errors.adult?.message && <Typography>{errors.adult.message}</Typography>}
            </Label>
            <Label>
              Child:
              <input type="number" min="0" {...register("child")} />
              {errors.child?.message && <Typography>{errors.child.message}</Typography>}
            </Label>
            <Label>
              Infant:
              <input type="number" min="0" {...register("infant")} />
              {errors.infant?.message && <Typography>{errors.infant.message}</Typography>}
            </Label>
            <input type="text" {...register("roomId")} hidden value={data.data.id} />
            {errors.roomId?.message && <Typography>{errors.roomId.message}</Typography>}
            <Button type="submit" disabled={isLoading}>Reserve</Button>
          </form>
        </div>}
    </div>
  );
}
