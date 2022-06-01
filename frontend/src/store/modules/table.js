import axios from "axios";
const state = () => ({
  items: [],
});
const mutations = {
  updateData(state, data) {
    state.items.splice(0, state.items.length);
    Object.assign(state.items, data);
  },
  newData(state, data) {
    if (state.items.length >= 10) {
      state.items.unshift(JSON.parse(data));
      state.items.pop();
    } else {
      state.items.unshift(JSON.parse(data));
    }
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
    } else {
      console.log(validator);
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
