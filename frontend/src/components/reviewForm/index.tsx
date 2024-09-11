import { CreateReviewInput, createReviewSchema } from "../../zod/review";
import { Rating } from "@smastrom/react-rating"
import Label from "../label";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Typography from "../typography";
import Button from "../button";
import '@smastrom/react-rating/style.css'

type Props = {
  callback: (payload: CreateReviewInput) => void;
  roomId: string;
}
function ReviewForm({ callback, roomId }: Props) {
  const { control, reset, register, handleSubmit, formState: { errors, isLoading } } = useForm<CreateReviewInput>({
    defaultValues: {
      rating: 0,
      comment: "",
      roomId: roomId ?? ""
    },
    resolver: zodResolver(createReviewSchema)
  })

  const onSubmit = handleSubmit((data) => {
    callback(data);
    reset();
  })
  return <form onSubmit={onSubmit} className="flex flex-col gap-2 w-full max-w-[350px] min-h-[100px] p-2 border-[1px] border-slate-200 rounded-md">
    <Label>
      Rating:
      <Controller
        name="rating"
        control={control}
        render={({ field }) =>
          <Rating className="max-w-[150px]" value={field.value} onChange={field.onChange} />
        }
      />
      {errors.rating?.message && <Typography>{errors.rating.message}</Typography>}
    </Label>
    <Label>
      Comment:
      <textarea {...register("comment")} className="border-[1px] rounded-md indent-2 border-slate-200 focus:outline-none focus:border-0 focus:ring-2 focus:ring-red-200" />
      {errors.comment?.message && <Typography>{errors.comment.message}</Typography>}
    </Label>
    <input hidden {...register("roomId")} />
    <Button type="submit" disabled={isLoading}>Add Comment</Button>
  </form>
}

export default ReviewForm;
