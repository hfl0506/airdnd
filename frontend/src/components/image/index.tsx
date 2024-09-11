import { ComponentProps, SyntheticEvent, useState } from "react";
import { cn } from "../../utils/cn";
import ExampleRoom from "../../assets/exampleRoom.webp";

interface Props extends ComponentProps<"img"> {}

function CustomImage({ src, className, alt, ...props }: Props) {
  const [isLoading, setIsLoading] = useState(true);
  const [hasError, setHasError] = useState(false);

  const handleImageLoad = () => {
    setIsLoading(false);
  };

  const handleImageError = (event: SyntheticEvent<HTMLImageElement, Event>) => {
    setIsLoading(false);
    setHasError(true);
    event.currentTarget.src = ExampleRoom;
  };

  return (
    <div className={cn("relative", className)}>
      {isLoading && (
        <div className="absolute inset-0 bg-gray-200 animate-pulse rounded" />
      )}
      <img
        src={hasError ? ExampleRoom : src}
        className={cn(
          "transition-opacity duration-500",
          isLoading ? "opacity-0" : "opacity-100",
          className
        )}
        alt={alt}
        loading="lazy"
        onLoad={handleImageLoad}
        onError={handleImageError}
        {...props}
      />
    </div>
  );
}

export default CustomImage;
