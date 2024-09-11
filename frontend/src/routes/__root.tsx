import * as React from "react";
import {
  Link,
  Outlet,
  createRootRouteWithContext,
} from "@tanstack/react-router";
import Navbar from "../components/navbar";
import { QueryClient } from "@tanstack/react-query";

const TanStackRouterDevtools =
  import.meta.env.NODE_ENV === "production"
    ? () => null
    : React.lazy(() =>
        import("@tanstack/router-devtools").then((res) => ({
          default: res.TanStackRouterDevtools,
        }))
      );

export const Route = createRootRouteWithContext<{ queryClient: QueryClient }>()(
  {
    component: RootComponent,
    notFoundComponent: NotFoundComponent,
  }
);

function RootComponent() {
  return (
    <React.Fragment>
      <Navbar />
      <hr />
      <Outlet />
      <React.Suspense>
        <TanStackRouterDevtools />
      </React.Suspense>
    </React.Fragment>
  );
}

function NotFoundComponent() {
  return (
    <div>
      <p>Route not found</p>
      <Link to="/" search={{ page: 1, limit: 10, tab_id: "" }}>
        Start Over
      </Link>
    </div>
  );
}
