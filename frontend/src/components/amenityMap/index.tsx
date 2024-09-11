import {
  AlarmSmoke,
  Baby,
  Bath,
  BriefcaseMedical,
  Cable,
  ChefHat,
  CircleParking,
  FireExtinguisher,
  FlameKindling,
  Grid2X2,
  Tv,
  Utensils,
  Waves,
  Wifi,
} from "lucide-react";

type IconLabelProps = {
  icon: React.ReactNode;
  text: string;
};

type AmenityMapProps = {
  amenity: string;
};

function IconLabel({ icon, text }: IconLabelProps) {
  return (
    <div className="flex items-center gap-4">
      {icon}
      <span>{text}</span>
    </div>
  );
}

function AmenityMap({ amenity }: AmenityMapProps) {
  const lowerKey = amenity.toLowerCase();
  let icon: React.ReactNode;
  if (lowerKey.includes("tv")) {
    icon = <Tv />;
  } else if (lowerKey.includes("cable")) {
    icon = <Cable />;
  } else if (lowerKey.includes("internet") || lowerKey.includes("wifi")) {
    icon = <Wifi />;
  } else if (lowerKey.includes("pool")) {
    icon = <Waves />;
  } else if (lowerKey.includes("breakfast")) {
    icon = <Utensils />;
  } else if (lowerKey.includes("kitchen")) {
    icon = <ChefHat />;
  } else if (lowerKey.includes("parking")) {
    icon = <CircleParking />;
  } else if (lowerKey.includes("fire") && !lowerKey.includes("extinguisher")) {
    icon = <FlameKindling />;
  } else if (lowerKey.includes("extinguisher")) {
    icon = <FireExtinguisher />;
  } else if (lowerKey.includes("kid") || lowerKey.includes("baby")) {
    icon = <Baby />;
  } else if (lowerKey.includes("smoke detector")) {
    icon = <AlarmSmoke />;
  } else if (lowerKey.includes("first aid")) {
    icon = <BriefcaseMedical />;
  } else if (lowerKey.includes("shampoo") || lowerKey.includes("bath")) {
    icon = <Bath />;
  } else {
    icon = <Grid2X2 />;
  }

  return <IconLabel icon={icon} text={amenity} />;
}

export default AmenityMap;
