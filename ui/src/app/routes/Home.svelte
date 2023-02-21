<script lang="ts">
  import { onMount } from "svelte";

  let count = 0;

  onMount(() => {
    const ih = setInterval(() => {
      count = count + 1;
    }, 1000);

    return () => clearInterval(ih);
  });

  const data = fetch("/api/data").then((resp) => resp.json());
</script>

<div class="w-full h-screen bg-black text-white grid place-items-center">
  <div class="flex flex-col gap-6 text-center">
    <h1 class="text-red-400 text-3xl">Welcome to Svelte SPA!</h1>
    <p class="max-w-2xl">
      This is a Svelte SPA. So everything on this page including the text and
      wiggly counter below are rendered using javascript on client-side.
    </p>
    <p class="animate-bounce text-xl">{count}</p>
    <div class="mt-6">
      <a
        class="rounded-md py-4 px-12 text-white bg-gradient-to-br from-red-500 to-cyan-600 p-1 hover:border font-medium"
        href="/">Go Home</a
      >
    </div>

    {#await data}
      <p>Waiting...</p>
    {:then resp}
      <p>{JSON.stringify(resp)}</p>
    {:catch error}
      <p>Error: {JSON.stringify(error)}</p>
    {/await}
  </div>
</div>
