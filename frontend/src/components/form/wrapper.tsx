import { ComponentProps } from "react";

interface Props extends ComponentProps<"div"> {}

function FormWrapper({ children }: Props) {
  return (
    <div className="w-full min-h-screen flex flex-col items-center mt-20">
      {children}
    </div>
  );
}

export default FormWrapper;
