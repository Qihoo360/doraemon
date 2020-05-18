import React, { Component } from 'react'
import { Modal, message } from 'antd'
import { Form, Input, TimePicker, InputNumber, DatePicker, TextArea, Row, Col, Checkbox, Button, AutoComplete } from 'antd'
import { formItemLayout, formItemLayoutWithOutLabel } from '@configs/const'

export default class HostList extends Component {
  componentDidMount() {
    this.props.onRef(this)
  }
  state = {
    visible: false,
    list: [],
  }
  updateValue(value) {
    this.setState({
      visible: true,
      list: value,
    })
  }
  handleOk = () => {
    this.setState({
      visible: false,
      list: [],
    })
  }
  render() {
    const { list } = this.state
    return (
      <Modal
        title="机器列表"
        visible={this.state.visible}
        onOk={this.handleOk}
        onCancel={this.handleOk}
        footer={[
          <Button key="submit" type="primary" onClick={this.handleOk}>
            确定
          </Button>,
        ]}
      >
        {
          list && list.map(item => (<p>{item}</p>))
        }
      </Modal>
    )
  }
}
