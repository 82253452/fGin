import {query as queryUsers, queryCurrent, queryMenu} from '../services/user';
import {getAuthority} from '../utils/authority';
export default {
  namespace: 'user',

  state: {
    list: [],
    currentUser: {},
    Menu: [],
  },

  effects: {
    *fetch(_, {call, put}) {
      const response = yield call(queryUsers);
      yield put({
        type: 'save',
        payload: response,
      });
    },
    *fetchCurrent(_, {call, put}) {
      const response = yield call(queryCurrent, {username: getAuthority()});
      yield put({
        type: 'saveCurrentUser',
        payload: response,
      });
    },
    *fetchMenuData(_, {call, put}) {
      const response = yield call(queryMenu, {username: getAuthority()});
      yield put({
        type: 'saveMenu',
        payload: response,
      });
    },
  },

  reducers: {
    save(state, action) {
      return {
        ...state,
        list: action.payload,
      };
    },
    saveCurrentUser(state, action) {
      return {
        ...state,
        currentUser: action.payload,
      };
    },
    changeNotifyCount(state, action) {
      return {
        ...state,
        currentUser: {
          ...state.currentUser,
          notifyCount: action.payload,
        },
      };
    },
    saveMenu(state, action) {
      return {
        ...state,
        Menu: action.payload,
      };
    },
  },
};
