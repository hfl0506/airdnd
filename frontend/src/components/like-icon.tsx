import { BookHeart, Heart } from "lucide-react";
import { ComponentProps } from "react";
import { cn } from "../utils/cn";
interface Props extends ComponentProps<"svg"> {
  isIncluded: boolean;
}
function LikeIcon({ onClick, isIncluded, className, ...rest }: Props) {
  const Comp = isIncluded ? BookHeart : Heart
  return <Comp onClick={onClick} className={cn("z-50 cursor-pointer absolute top-1 right-1 -translate-x-1 translate-y-1 w-6 h-6 text-white hover:text-red-300", className)} {...rest} />
}

export default LikeIcon;
