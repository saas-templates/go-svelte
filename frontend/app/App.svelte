<script lang="ts">
  import Router from "svelte-spa-router";
  import { wrap } from "svelte-spa-router/wrap";
  import Home from "./pages/Home.svelte";

  import NotFound from "./pages/NotFound.svelte";

  const routes = {
    "/": Home,
    "*": NotFound,
  };
</script>

<Router {routes} />
