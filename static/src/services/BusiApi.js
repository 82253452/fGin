import request from '../utils/request';
import {stringify} from 'qs';

export async function queryUser(params) {
  return request('/sys/queryUser');
}
export async function queryRole(params) {
  return request('/sys/queryRole');
}
export async function queryTreeMenu(params) {
  return request(`/sys/queryTreeMenu?${stringify(params)}`);
}
export async function querySelectedKeys(params) {
  return request(`/sys/querySelectedKeys?${stringify(params)}`);
}
export async function addRole(params) {
  return request(`/sys/addRole?${stringify(params)}`);
}
export async function updateRoleMenu(params) {
  return request(`/sys/updateRoleMenu?${stringify(params)}`);
}
export async function upDateRole(params) {
  return request(`/sys/upDateRole?${stringify(params)}`);
}
export async function delRole(params) {
  return request(`/sys/delRole?${stringify(params)}`);
}
export async function queryByRoleId(params) {
  return request(`/sys/queryByRoleId?${stringify(params)}`);
}

