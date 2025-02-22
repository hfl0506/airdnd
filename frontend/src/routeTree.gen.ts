/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as IndexImport } from './routes/index'
import { Route as WishlistIndexImport } from './routes/wishlist/index'
import { Route as SignupIndexImport } from './routes/signup/index'
import { Route as LoginIndexImport } from './routes/login/index'
import { Route as BookingsIndexImport } from './routes/bookings/index'
import { Route as RoomsRoomIdIndexImport } from './routes/rooms/$roomId/index'

// Create Virtual Routes

const AboutLazyImport = createFileRoute('/about')()

// Create/Update Routes

const AboutLazyRoute = AboutLazyImport.update({
  path: '/about',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/about.lazy').then((d) => d.Route))

const IndexRoute = IndexImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const WishlistIndexRoute = WishlistIndexImport.update({
  path: '/wishlist/',
  getParentRoute: () => rootRoute,
} as any)

const SignupIndexRoute = SignupIndexImport.update({
  path: '/signup/',
  getParentRoute: () => rootRoute,
} as any)

const LoginIndexRoute = LoginIndexImport.update({
  path: '/login/',
  getParentRoute: () => rootRoute,
} as any)

const BookingsIndexRoute = BookingsIndexImport.update({
  path: '/bookings/',
  getParentRoute: () => rootRoute,
} as any)

const RoomsRoomIdIndexRoute = RoomsRoomIdIndexImport.update({
  path: '/rooms/$roomId/',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutLazyImport
      parentRoute: typeof rootRoute
    }
    '/bookings/': {
      id: '/bookings/'
      path: '/bookings'
      fullPath: '/bookings'
      preLoaderRoute: typeof BookingsIndexImport
      parentRoute: typeof rootRoute
    }
    '/login/': {
      id: '/login/'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginIndexImport
      parentRoute: typeof rootRoute
    }
    '/signup/': {
      id: '/signup/'
      path: '/signup'
      fullPath: '/signup'
      preLoaderRoute: typeof SignupIndexImport
      parentRoute: typeof rootRoute
    }
    '/wishlist/': {
      id: '/wishlist/'
      path: '/wishlist'
      fullPath: '/wishlist'
      preLoaderRoute: typeof WishlistIndexImport
      parentRoute: typeof rootRoute
    }
    '/rooms/$roomId/': {
      id: '/rooms/$roomId/'
      path: '/rooms/$roomId'
      fullPath: '/rooms/$roomId'
      preLoaderRoute: typeof RoomsRoomIdIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  IndexRoute,
  AboutLazyRoute,
  BookingsIndexRoute,
  LoginIndexRoute,
  SignupIndexRoute,
  WishlistIndexRoute,
  RoomsRoomIdIndexRoute,
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/about",
        "/bookings/",
        "/login/",
        "/signup/",
        "/wishlist/",
        "/rooms/$roomId/"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/about": {
      "filePath": "about.lazy.tsx"
    },
    "/bookings/": {
      "filePath": "bookings/index.tsx"
    },
    "/login/": {
      "filePath": "login/index.tsx"
    },
    "/signup/": {
      "filePath": "signup/index.tsx"
    },
    "/wishlist/": {
      "filePath": "wishlist/index.tsx"
    },
    "/rooms/$roomId/": {
      "filePath": "rooms/$roomId/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
