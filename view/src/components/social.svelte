<script>
  import Icon from "./icon.svelte"

  export let icon
  export let link
</script>

<a href="{link}" target="_blank" class="social">
  <Icon name="{icon}" style="social-icon" />
</a>

<style>
  .social {
    text-decoration: none;

    margin-top: 35px;

    padding-left: 25px;
    padding-right: 25px;
    padding-top: 15px;
    padding-bottom: 15px;

    display: inline-flex;
    align-items: center;

    transition-duration: 0.3s;
  }

  .social:hover {
    opacity: 0.7;
  }

  :global([style="social-icon"]) {
      height: 45px;
    }

  @media only screen and (max-width: 540px) {
    :global([style="social-icon"]) {
      height: 30px;
    }
  }

  @media only screen and (max-height: 600px) {
    .social {
      margin-top: 0;
    }
  }

  @media only screen and (max-width: 330px) {
    .social {
      padding-left: 20px;
      padding-right: 20px;
    }
  }
</style>