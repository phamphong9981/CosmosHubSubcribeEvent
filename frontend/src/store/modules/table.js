import axios from "axios";

const state = () => ({
  items: [],
});
const mutations = {
  updateData(state, data) {
    Object.assign(state.items, data);
  },
  newData(state, data) {
    state.items.unshift(JSON.parse(data));
    state.items.pop();
  },
};

const actions = {
  async getData(context) {
    const response = await axios.get("http://localhost:8088/unbond/all");
    const data = [];
    response.data.map((item) => {
      data.push(JSON.parse(item));
    });
    context.commit("updateData", data);
  },
  async getDataValidator(context, [validator]) {
    if (validator) {
      const response = await axios.get(
        "http://localhost:8088/unbond/" + validator
      );
      const data = [];
      response.data.map((item) => {
        data.push(JSON.parse(item));
      });
      context.commit("updateData", data);
    }else{
      console.log("Break point");
      context.commit("updateData", []);
    }
  },
};

export default {
  namespaced: true,
  state,
  actions,
  mutations,
};
