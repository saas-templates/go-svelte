<script lang="ts">
  import Router from "svelte-spa-router";
  import { wrap } from "svelte-spa-router/wrap";

  import Home from "./routes/Home.svelte";
  import NotFound from "./routes/NotFound.svelte";

  const routes = {
    "/": Home,
    "*": NotFound,
  };
</script>

<div class="bg-black">
  <Router {routes} />
</div>
