/**
 * Created by admin on 2018/6/8.
 */
import {Form, Modal} from "antd";
const CreateForm = Form.create({
  mapPropsToFields(props) {
    var obj = {};
    Object.keys(props.model).map((key) => {
      obj[key] = Form.createFormField({
        value: props.model[key],
      });
    });
    return obj;
  }
})(props => {
  const {modelVisible, form, creatOK, cancelHandle, title, renderForm} = props;
  const okHandle = () => {
    form.validateFields((err, fieldsValue) => {
      if (err) return;
      form.resetFields();
      creatOK(fieldsValue);
    });
  };
  return (
    <Modal
      title={title}
      visible={modelVisible}
      onOk={okHandle}
      onCancel={cancelHandle}
    >
      {renderForm(form)}
    </Modal>
  );
});
export default CreateForm
