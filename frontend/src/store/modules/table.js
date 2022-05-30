const state = () => ({
  items: [],
});
const mutations = {
  updateData(state, data) {
    state.items.pop();
    state.items=data;
  },
};

const actions = {};

export default {
  namespaced: true,
  state,
  actions,
  mutations,
};
