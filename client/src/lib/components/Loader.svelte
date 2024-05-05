<script lang="ts">
  export let message: string
  export let show: boolean
</script>

<main class="w-full h-full absolute center top-0 left-0 z-50 flex-col gap-y-4">
  <div class="mb-5 -ml-1">
    <span>{message}</span>
  </div>
  {#if show}
    <div class="loader" />
  {/if}
</main>

<style>
    /* HTML: <div class="loader"></div> */
    .loader {
        height: 15px;
        aspect-ratio: 4;
        --_g: no-repeat radial-gradient(farthest-side,#ffffff 90%,#0000);
        background:
                var(--_g) left,
                var(--_g) right;
        background-size: 25% 100%;
        display: grid;
    }
    .loader:before,
    .loader:after {
        content: "";
        height: inherit;
        aspect-ratio: 1;
        grid-area: 1/1;
        margin: auto;
        border-radius: 50%;
        transform-origin: -100% 50%;
        background: #ffffff;
        animation: l49 1s infinite linear;
    }
    .loader:after {
        transform-origin: 200% 50%;
        --s:-1;
        animation-delay: -.5s;
    }

    @keyframes l49 {
        58%,
        100% {transform: rotate(calc(var(--s,1)*1turn))}
    }
</style>
