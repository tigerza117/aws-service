<script>
  import { Route } from "svelte-navigator";
  import PublicRouteGuard from "./Public.svelte";

  export let path;
</script>

<Route {path} let:params let:location let:navigate>
  <PublicRouteGuard>
    <slot {params} {location} {navigate} />
  </PublicRouteGuard>
</Route>
