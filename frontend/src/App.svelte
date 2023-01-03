<script lang="ts">
  import Router from "svelte-spa-router";
  import { wrap } from "svelte-spa-router/wrap";
  import Landing from "./pages/Landing.svelte";

  import NotFound from "./pages/NotFound.svelte";

  const routes = {
    "/": Landing,
    "*": NotFound,
  };
</script>

<Router {routes} />
