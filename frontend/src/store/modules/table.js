import axios from "axios";

const state = () => ({
  items: [],
});
const mutations = {
  updateData(state, data) {
    Object.assign(state.items,data)
  },
};

const actions = {
  async getData(context) {
    const response = await axios.get("http://localhost:8088/unbond/all");
    const data=[]
    response.data.map(item=>{
      console.log(item);
      data.push(JSON.parse(item))
    })
    context.commit("updateData",data);
  },
  async streamRealData(){

  }
};

export default {
  namespaced: true,
  state,
  actions,
  mutations,
};
