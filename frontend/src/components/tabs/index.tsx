import { House } from "lucide-react";
import { cn } from "../../utils/cn";
import { useEffect, useRef, useState } from "react";

type Props = {
  tabs: Array<string>;
  activeTab: string;
  onClick: (tab: string) => void;
};

function Tabs({ tabs, activeTab, onClick }: Props) {
  const scrollRef = useRef<HTMLUListElement>(null);
  const [isScrolled, setIsScrolled] = useState(false);
  const [canScrollRight, setCanScrollRight] = useState(false);
  const DISTANCE = 200;

  useEffect(() => {
    const handleScroll = () => {
      const el = scrollRef.current;

      if (el !== null) {
        setIsScrolled(el.scrollLeft! > 0);
        setCanScrollRight(el.scrollLeft + el.clientWidth < el.scrollWidth);
      }
    };

    handleScroll();
    if (scrollRef.current !== null) {
      scrollRef?.current.addEventListener("scroll", handleScroll);
    }
    return () => {
      if (scrollRef.current !== null) {
        scrollRef.current.removeEventListener("scroll", handleScroll);
      }
    };
  }, []);

  const scrollLeft = () => {
    if (scrollRef.current !== null) {
      scrollRef.current.scrollBy({ left: -DISTANCE, behavior: "smooth" });
    }
  };

  const scrollRight = () => {
    if (scrollRef.current !== null) {
      scrollRef.current.scrollBy({ left: DISTANCE, behavior: "smooth" });
    }
  };

  return (
    <div className="relative w-full max-w-[1600px] mx-auto px-4">
      {isScrolled && (
        <button
          className="absolute left-6 top-1/2 -translate-y-1/2 bg-slate-100 rounded-full w-10 h-10"
          onClick={scrollLeft}
        >
          &larr;
        </button>
      )}
      <ul
        ref={scrollRef}
        className="w-full h-14 px-2 bg-slate-200 rounded-md flex items-center gap-2 overflow-x-auto overflow-y-hidden scrollbar-hide"
      >
        {tabs?.length > 0
          ? tabs.map((tab) => (
              <button
                key={tab}
                onClick={() => onClick(tab)}
                className={cn(
                  "flex items-center gap-2 px-4 py-2 rounded-md",
                  activeTab === tab
                    ? "bg-white text-black"
                    : "bg-transparent text-slate-500"
                )}
              >
                <House className="w-5 h-5 flex-shrink-0" />
                <span className="truncate">{tab}</span>
              </button>
            ))
          : null}
      </ul>
      {canScrollRight && (
        <button
          className="absolute right-6 top-1/2 -translate-y-1/2 bg-slate-100 rounded-full w-10 h-10"
          onClick={scrollRight}
        >
          &rarr;
        </button>
      )}
    </div>
  );
}

export default Tabs;
