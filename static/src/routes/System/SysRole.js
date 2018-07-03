import React, {Component, Fragment} from "react";
import {connect} from "dva";
import {Button, Card, Form, Input, Modal, Tree, Divider} from "antd";
import StandardTable from "components/StandardTable";
import PageHeaderLayout from "../../layouts/PageHeaderLayout";
import CreateForm from "../../cusComponents/FormModel/CreateForm";
import styles from "./TableList.less";
const FormItem = Form.Item;

const AuthForm = Form.create()(props => {
  const {show, okHandle, cancelHandle, treeData, expandedKeys, onTreeCheck} = props;
  const TreeNode = Tree.TreeNode;
  const onCheck = (checkedKeys) => {
    onTreeCheck(checkedKeys);
  };
  const renderTreeNodes = (data) => {
    return data == null ? '' : data.map((item) => {
      if (item.children) {
        return (
          <TreeNode title={item.title} key={item.key} dataRef={item}>
            {renderTreeNodes(item.children)}
          </TreeNode>
        );
      }
      return <TreeNode {...item} />;
    });
  };
  return (
    <Modal
      title="分配权限"
      visible={show}
      onOk={okHandle}
      onCancel={() => cancelHandle()}
    >
      {treeData.length
        ? <Tree
          checkable
          checkedKeys={expandedKeys}
          onCheck={onCheck}
        >
          {renderTreeNodes(treeData)}
        </Tree>
        : 'loading tree'}

    </Modal>
  );
});

@connect(({SysModel,}) => ({
  SysModel,
}))
@Form.create()
export default class SysRole extends Component {
  state = {
    modalVisible: false,
    selectedRows: [],
    showAuthority: false,
    createShow: false,
  };

  componentDidMount() {
    const {dispatch} = this.props;
    dispatch({
      type: 'SysModel/queryRole',
    });
    this.props.dispatch({
      type: 'SysModel/queryTreeMenu',
      payload: {
        roleId: "2",
      },
    });
  }

  showAuthority = (roleId) => {
    this.props.dispatch({
      type: 'SysModel/querySelectedKeys',
      payload: {
        roleId: roleId,
      },
    });
    this.setState({
      showAuthority: true,
    });
  };
  okHandle = () => {
    this.setState({
      showAuthority: false,
    });
  };
  showTogle = () => {
    this.props.dispatch({
      type: 'SysModel/clearCurrentRole',
    });
    this.setState({
      createShow: !this.state.createShow,
    });
  };
  creatOK = (files) => {
    this.props.dispatch({
      type: 'SysModel/addRole',
      payload: {
        ...this.props.SysModel.currentRole,
        ...files
      },
    });
    this.setState({
      createShow: false,
    });
  };
  onTreeCheck = (selectedKeys) => {
    this.props.dispatch({
      type: 'SysModel/upSelectedKeys',
      payload: {
        selectedKeys: selectedKeys,
      },
    });
  };
  renderCForm = (form) => {
    return (<div>
      <FormItem labelCol={{span: 5}} wrapperCol={{span: 15}} label="角色名称">
        {form.getFieldDecorator('RoleName', {
          rules: [{required: true, message: 'Please input some description...'}],
        })(<Input placeholder="角色名称"/>)}
      </FormItem>
      <FormItem labelCol={{span: 5}} wrapperCol={{span: 15}} label="备注">
        {form.getFieldDecorator('Remark')(<Input placeholder="备注"/>)}
      </FormItem>
    </div>)
  };
  updateRole = (roleId) => {
    this.showTogle();
    this.props.dispatch({
      type: 'SysModel/queryByRoleId',
      payload: {
        roleId: roleId,
      },
    });
  };
  delRole = () => {
    this.props.dispatch({
      type: 'SysModel/delRole',
      payload: {
        roleIds: JSON.stringify(this.state.selectedRows),
      },
    });
  };
  handleSelectRows = (selectedRows) => {
    this.setState({
      selectedRows: selectedRows,
    });
  };
  cancelHandle = () => {
    this.setState({
      showAuthority: false,
    });
  };


  render() {

    const {SysModel: {role, treeMenu, selectedKeys, currentRole}} = this.props;
    const {selectedRows, createShow, showAuthority,} = this.state;
    const CreateMethods = {
      creatOK: this.creatOK,
      cancelHandle: this.showTogle,
      modelVisible: createShow,
      renderForm: this.renderCForm,
      model: currentRole
    };
    const columns = [
      {
        title: '角色名称',
        dataIndex: 'RoleName',
      },
      {
        title: '操作',
        render: (text, record) => {
          return ( <Fragment>
            <a onClick={() => this.showAuthority(record.RoleId)}>分配权限</a>
            <Divider type="vertical"/>
            <a onClick={() => this.updateRole(record.RoleId)}>更新</a>
          </Fragment>)
        },
      },
    ];
    return (
      <PageHeaderLayout title="查询表格">
        <Card bordered={false}>
          <div className={styles.tableList}>
            <div className={styles.tableListOperator}>
              <Button icon="plus" type="primary" onClick={() => this.showTogle()}>
                新建
              </Button>
              <Button icon="plus" type="primary" onClick={() => this.delRole()}>
                删除
              </Button>
            </div>
            <StandardTable
              selectedRows={selectedRows}
              // loading={loading}
              data={role}
              columns={columns}
              rowKey="RoleId"
              onSelectRow={this.handleSelectRows}
              // onChange={this.handleStandardTableChange}
            />
          </div>
        </Card>
        <AuthForm show={showAuthority} treeData={treeMenu} cancelHandle={this.cancelHandle}
                  okHandle={this.okHandle}
                  expandedKeys={selectedKeys}
                  onTreeCheck={this.onTreeCheck}/>
        <CreateForm {...CreateMethods} />
      </PageHeaderLayout>
    );
  }
}
