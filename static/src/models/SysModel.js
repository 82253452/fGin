/**
 * Created by admin on 2018/6/6.
 */
import {
  addRole,
  delRole,
  queryByRoleId,
  queryRole,
  querySelectedKeys,
  queryTreeMenu,
  queryUser,
  upDateRole,
  updateRoleMenu
} from "../services/BusiApi";

export default {
  namespace: 'SysModel',

  state: {
    user: [],
    role: [],
    treeMenu: [],
    selectedKeys: [],
    currentRole: {}
  },

  effects: {
    *queryUser({payload}, {call, put}) {
      const response = yield call(queryUser, payload);
      yield put({
        type: 'saveUser',
        payload: response,
      });
    },
    *queryRole({payload, callback}, {call, put}) {
      const response = yield call(queryRole, payload);
      yield put({
        type: 'saveRole',
        payload: response,
      });
    },
    *queryTreeMenu({payload, callback}, {call, put}) {
      const response = yield call(queryTreeMenu, payload);
      yield put({
        type: 'saveTreeMenu',
        payload: response,
      });
    },
    *querySelectedKeys({payload, callback}, {call, put}) {
      const response = yield call(querySelectedKeys, payload);
      yield put({
        type: 'saveSelectedKey',
        payload: response,
      });
    },
    *addRole({payload, callback}, {call, put}) {
      yield call(addRole, payload);
      yield put({
        type: 'queryRole',
      });
    },
    *updateRoleMenu({payload, callback}, {call, put}) {
      yield call(updateRoleMenu, payload);
      yield put({
        type: 'queryRole',
      });
    },
    *upSelectedKeys({payload, callback}, {call, put}) {
      yield put({
        type: 'upSelectedKeysReducer',
        payload: payload,
      });
    },
    *upDateRole({payload, callback}, {call, put}) {
      const response = yield call(upDateRole, payload);
    },
    *delRole({payload, callback}, {call, put}) {
      yield call(delRole, payload);
      yield put({
        type: 'queryRole',
      });
    },
    *queryByRoleId({payload, callback}, {call, put}) {
      const response = yield call(queryByRoleId, payload);
      yield put({
        type: 'saveRoleByRoleId',
        payload: response,
      });
    }, *clearCurrentRole({payload, callback}, {call, put}) {
      yield put({
        type: 'saveRoleByRoleId',
        payload: {},
      });
    },
  },

  reducers: {
    saveUser(state, action) {
      return {
        ...state,
        user: action.payload,
      };
    }, saveRole(state, action) {
      return {
        ...state,
        role: action.payload,
      };
    }, saveTreeMenu(state, action) {
      return {
        ...state,
        treeMenu: action.payload,
      };
    }, saveSelectedKey(state, action) {
      return {
        ...state,
        selectedKeys: action.payload,
      };
    },
    upSelectedKeysReducer(state, action) {
      return {
        ...state,
        selectedKeys: action.payload.selectedKeys,
      };
    },
    saveRoleByRoleId(state, action) {
      return {
        ...state,
        currentRole: action.payload,
      };
    },
  },
};
