<template>
  <v-row justify="center" class="search-bar"
    ><v-col cols="8">
      <form action="" method="get">
        <input type="text" placeholder="Search by validator address..." />
        <button type="submit">
          <i class="fa-solid fa-magnifying-glass"></i>
        </button>
      </form> </v-col
  ></v-row>
</template>

<script>
export default {};
</script>

<style>
form {
  background: white;
  border-radius: 40px;
}
input[type="text"] {
  height: 57px;
  border: none;
  outline: none;
  color: black;
  width: 95%;
  padding: 15px;
}
button i{
  color: gray;
}
::placeholder {
  color: gray;
}
.search-bar {
  margin-top: 20px;
}
</style>