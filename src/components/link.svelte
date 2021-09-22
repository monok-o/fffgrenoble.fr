<script>
  import Icon from "./icon.svelte"

  export let content
  export let link
</script>

<a href="{link}" target="_blank" class="link">
  <Icon name="link" style="link-icon" />
  <span>{content}</span>
</a>

<style>
  .link {
    text-decoration: none;

    background-color: var(--color-dark);
    border-radius: 50px;

    margin-top: 35px;

    padding-left: 25px;
    padding-right: 25px;
    padding-top: 15px;
    padding-bottom: 15px;

    display: inline-flex;
    align-items: center;

    transition-duration: 0.3s;
  }

  .link:hover {
    opacity: 0.7;
  }

  :global([style="link-icon"]) {
    margin-right:25px;
  }

  .link span {
    color: white;
    font-size: 30px;
    font-weight: 700;
  }

  @media only screen and (max-width: 540px) {
    .link {
      margin-top: 25px;

      padding-left: 15px;
      padding-right: 15px;
      padding-top: 5px;
      padding-bottom: 5px;
    }

    :global([style="link-icon"]) {
      margin-right:10px;
      width: 18px;
    }

    .link span {
      font-size: 20px;
    }
  }
</style>