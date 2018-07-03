import React, {Component, Fragment} from "react";
import {connect} from "dva";
import {Card, Form, Input, Modal} from "antd";
import StandardTable from "components/StandardTable";
import PageHeaderLayout from "../../layouts/PageHeaderLayout";
import styles from "./TableList.less";
const FormItem = Form.Item;


const CreateForm = Form.create()(props => {
  const {modalVisible, form, handleAdd, handleModalVisible} = props;
  const okHandle = () => {
    form.validateFields((err, fieldsValue) => {
      if (err) return;
      form.resetFields();
      handleAdd(fieldsValue);
    });
  };
  return (
    <Modal
      title="新建规则"
      visible={modalVisible}
      onOk={okHandle}
      onCancel={() => handleModalVisible()}
    >
      <FormItem labelCol={{span: 5}} wrapperCol={{span: 15}} label="描述">
        {form.getFieldDecorator('desc', {
          rules: [{required: true, message: 'Please input some description...'}],
        })(<Input placeholder="请输入"/>)}
      </FormItem>
    </Modal>
  );
});

@connect(({SysModel,}) => ({
  SysModel,
}))
@Form.create()
export default class SysUser extends Component {

  state = {
    modalVisible: false,
    expandForm: false,
    selectedRows: [],
    formValues: {},
  };

  componentDidMount() {
    const {dispatch} = this.props;
    dispatch({
      type: 'SysModel/queryUser',
    });
  }
  renderForm() {
    return this.state.expandForm ? this.renderAdvancedForm() : this.renderSimpleForm();
  }

  render() {

    const {SysModel: {user}} = this.props;
    const {selectedRows, modalVisible} = this.state;
    const parentMethods = {
      handleAdd: this.handleAdd,
      handleModalVisible: this.handleModalVisible,
    };

    const columns = [
      {
        title: '姓名',
        dataIndex: 'Username',
      },
      {
        title: '手机号',
        dataIndex: 'Mobile',
      }, {
        title: '操作',
        render: () => (
          <Fragment>
            <a href="">更新</a>
          </Fragment>
        ),
      },
    ];
    return (
      <PageHeaderLayout title="查询表格">
        <Card bordered={false}>
          <div className={styles.tableList}>
            <StandardTable
              selectedRows={selectedRows}
              // loading={loading}
              data={user}
              columns={columns}
              // onSelectRow={this.handleSelectRows}
              // onChange={this.handleStandardTableChange}
            />
          </div>
        </Card>
        <CreateForm {...parentMethods} modalVisible={modalVisible}/>
      </PageHeaderLayout>
    );
  }
}
