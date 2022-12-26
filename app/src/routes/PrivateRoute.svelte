<script>
  import { Route } from "svelte-navigator";
  import PrivateRouteGuard from "./Private.svelte";

  export let path;
</script>

<Route {path} let:params let:location let:navigate>
  <PrivateRouteGuard>
    <slot {params} {location} {navigate} />
  </PrivateRouteGuard>
</Route>
