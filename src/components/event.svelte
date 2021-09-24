<script>
  import Separator from "./separator.svelte"

  export let title
  export let date
  export let time
  export let place
  export let link
</script>

<div>
  <Separator />
  <h3 class="head text title">{title}</h3>
  <div class="dateplace text">
    <p>{date} {#if time!=""}
      - {time}
    {/if}</p>
    {#if place}
      <p>{place}</p>
    {/if}
  </div>
  <a href={link} target="_blank" class="text link">Plus d'infos</a>
</div>

<style>
.title {
  font-size: 30px;
  font-weight: 700;
}

.dateplace {
    display: flex;
    justify-content: space-between;
}

.dateplace p {
    margin-bottom: 5px;
}

.text {
    color:white;
}

.link {
    margin-bottom:15px;
    display:inline-block;
}

@media only screen and (max-width: 540px) {
  .title {
    font-size: 20px;
  }
}

@media only screen and (max-width: 330px) {
  .title {
    font-size: 15px;
  }
}
</style>