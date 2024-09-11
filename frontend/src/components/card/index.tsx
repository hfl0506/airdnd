import { PropsWithChildren } from "react";
import { cn } from "../../utils/cn";

interface Props extends PropsWithChildren {
  onClick: () => void;
}

function Card({ children, onClick }: Props) {
  return (
    <div
      className={cn(
        "relative w-[350px] h-[400px] rounded-md border-slate-300 flex flex-col gap-4 cursor-pointer z-[1] md:w-[250px]"
      )}
      onClick={onClick}
    >
      {children}
    </div>
  );
}

export default Card;
