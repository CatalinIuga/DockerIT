<script>
  let container;
  window.go.main.App.GetContainerList().then((containers) => {
    container = containers;
    console.log(containers);
  });

  let a = "";
  function ping() {
    window.go.main.App.Ping().then((res) => {
      a = res;
    });
  }
</script>

<main>
  <button on:click={ping}>Ping!</button>
  <p>{a}</p>
</main>

<style>
</style>
