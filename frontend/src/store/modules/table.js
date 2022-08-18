import axios from "axios";
const state = () => ({
  items: [],
  delegator: "",
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
  viewMore(state,data){
    if(data){
      data.map((item) => {
        state.items.push(JSON.parse(item));
      });
    }
  },
  updateDelegator(state,delegator){
    state.delegator=delegator
    console.log(delegator);
  }
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
  async getDataDelegator(context, [delegator]) {
    if (delegator) {
      const response = await axios.get(
        "http://localhost:8088/unbond/" + delegator
      );
      const data = [];
      response.data.map((item) => {
        data.push(JSON.parse(item));
      });
      context.commit("updateData", data);
    } else {
      console.log(delegator);
      context.commit("updateData", []);
    }
    context.commit("updateDelegator",delegator)
  },
  async viewMoreAll(context){
    const response = await axios.get("http://localhost:8088/unbond/all?view_more_offset="+context.state.items.length);
    context.commit("viewMore", response.data);
  },
  async viewMoreByDelegator(context){
    if(state.delegator){
      const response = await axios.get("http://localhost:8088/unbond/"+context.state.delegator+"?view_more_offset="+context.state.items.length);
      context.commit("viewMore", response.data);
    }
  }
};

export default {
  namespaced: true,
  state,
  actions,
  mutations,
};
