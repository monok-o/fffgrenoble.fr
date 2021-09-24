<script>
export let title
export let date
export let time
export let place
export let link
</script>

<div>
    <h3 class="head text">{title}</h3>
    <div class="dateplace text">
        <p>{date} {#if time!=""}
            - {time}
        {/if}</p>
        {#if place}
        <p>{place}</p>
        {/if}
    </div>
    <a href={link} class="text link">Plus d'infos</a>
</div>

<style>
h3 {
    font-weight: 700;
    margin-bottom: 0;
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
</style>