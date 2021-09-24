<script>
  export let side
</script>

{#if side == "topleft" }
  <div class="wave topleft">
    <svg viewBox="0 0 1000 1000" preserveAspectRatio="xMinYMin meet">
      <path d="M0 0H1000C1000 0 393.193 78.5 337.5 423.5C290.195 716.538 0 1000 0 1000V0Z" fill="#36DC6C"/>
    </svg>
  </div>
{:else if side == "bottomright" }
  <div class="wave bottomright">
    <svg viewBox="0 0 1000 1000" preserveAspectRatio="xMinYMin meet">
      <path d="M1001 1000L1.00012 1000C1.00012 1000 -17 898 527.5 648.5C845.453 502.808 1001 0 1001 0L1001 1000Z" fill="#36DC6C"/>
    </svg>
  </div>
{/if}



<style>
  svg {
    display: inline-block;
    position: absolute;
    top: 0;
    left: 0;
  }
  .wave {
    position: fixed;

    width: 550px;
    height: 550px;
    z-index: -1;
  }

  @media only screen and (max-width: 1400px) {
    .wave {
      width: 400px;
      height: 400px;
    }
  }

  @media only screen and (max-width: 1110px) {
    .wave {
      width: 300px;
      height: 300px;
    }
  }

  @media only screen and (max-width: 800px) {
    .wave {
      width: 200px;
      height: 200px;
    }
  }

  @media only screen and (max-width: 370px) {
    .wave {
      width: 150px;
      height: 150px;
    }
  }

  @media only screen and (max-width: 370px) {
    .wave {
      width: 150px;
      height: 150px;
    }
  }

  @media only screen and (max-width: 330px) {
    .wave {
      display: none;
    }
  }

  @media only screen and (max-height: 600px) {
    .wave {
      display: none;
    }
  }

  .topleft {
    left: 0;
    top: 0;
  }

  .bottomright {
    right: 0;
    bottom: 0;
  }
</style>