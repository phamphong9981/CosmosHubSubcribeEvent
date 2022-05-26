<template>
  <div class="search-bar">
      <form action="" method="get">
          <input type="text" placeholder="Search...">
      </form>
  </div>
</template>

<script>
export default {

}
</script>

<style>
input[type="text"]{
    background: white;
    width: 80%;
    border: none;
    outline: none;
    color: black;
}
::placeholder{
    color: gray;
}
.search-bar{
    margin-top: 20px;

}
</style>