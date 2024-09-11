import { ComponentProps } from "react";
import { cn } from "../../utils/cn";

interface Props extends ComponentProps<"label"> { }

function Label({ children, className, ...rest }: Props) { return <label className={cn(className, "flex flex-col gap-2")} {...rest}>{children}</label> }

export default Label;
