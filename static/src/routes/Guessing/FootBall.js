import React, {PureComponent} from "react";
import {connect} from "dva";
import {Input, Radio} from "antd";

import PageHeaderLayout from "../../layouts/PageHeaderLayout";

const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;
const {Search} = Input;

@connect(({list, loading}) => ({
  list,
  loading: loading.models.list,
}))
export default class BasicList extends PureComponent {
  componentDidMount() {
    this.props.dispatch({
      type: 'footBall/fetch',
      payload: {
        count: 5,
      },
    });
  }

  render() {
    const {list: {list}, loading} = this.props;

    return (
      <div>123213123</div>
    );
  }
}
