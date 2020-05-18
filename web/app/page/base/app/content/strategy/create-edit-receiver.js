import React, { Component } from 'react'
import { Modal, message, Form, Input, TimePicker, InputNumber, Select } from 'antd'
import { formItemLayout } from '@configs/const'
import moment from 'moment'
import { DInput } from '@components/input/input'

const { Option } = Select;
@Form.create({})
export default class CreateEditReceiver extends Component {
  constructor(props) {
    super(props)
    this.props.OnRef(this)
  }
  state = {
    id: 0,
    mode: 'create', // create or edit
    method: 'HOOK',
    visible: false,
    groupState: {
      user: true,
      duty_group: true,
      group: true,
    },
  }
  componentDidMount() {
    const { form } = this.props
    form.setFieldsValue()
  }
  selectMethod = (value) => {
    // console.log(value)
    this.setState({
      method: value,
    })
  }
  updateValue(value) {
    const { form } = this.props
    const { mode, id, ...data } = value
    // console.log(mode)
    // console.log(data.method.substr(0,4))
    var method = "LANXIN"
    if (mode === 'edit') {
      if (data.method.length >= 4 && data.method.substr(0, 4) === "HOOK") {
        method = "HOOK"
        data.hookurl = data.method.substr(5)
        data.method = "HOOK"
      } else {
        method = data.method
      }
    }
    // console.log(data)
    this.setState({
      id,
      mode,
      method: method,
      visible: true,
      groupState: {
        user: true,
        duty_group: true,
        group: true,
      },
    })
    form.resetFields()
    mode === 'edit' && form.setFieldsValue(this.unFormatValue(data))
  }
  formatValue(values) {
    const { ...value } = values
    value.start_time = moment(value.start_time).format('HH:mm')
    value.end_time = moment(value.end_time).format('HH:mm')
    return value
  }
  unFormatValue(values) {
    const { ...value } = values
    value.start_time = moment(value.start_time, 'HH:mm')
    value.end_time = moment(value.end_time, 'HH:mm')
    const groupValue = !(values.user || values.duty_group || values.duty_group === 0 || values.group || values.group === 0)
    this.setState({
      groupState: {
        user: groupValue,
        duty_group: groupValue,
        group: groupValue,
      },
    })
    // console.log(value)
    return value
  }
  handleOk = () => {
    this.props.form.validateFields(async (err, values) => {
      if (!err) {
        const { id, mode, method } = this.state
        if (method === "HOOK") {
          values.method = values.method + " " + values.hookurl
        }
        const resultSuccess = await this.props.onSubmit({ id, mode, ...this.formatValue(values) })
        if (resultSuccess) {
          message.success(mode === 'edit' ? '编辑成功' : '添加成功')
          this.setState({
            method: "HOOK",
            visible: false,
          })
        }
      }
    })
  }
  handleCancel = () => {
    const { form } = this.props
    this.setState({
      method: "HOOK",
      visible: false,
    })
    // form.resetFields()
  }
  startTimeChange = (value) => {
    const { form } = this.props
    const endTime = form.getFieldValue('end_time')
    if (endTime) {
      form.validateFields(['end_time'], { force: true })
    }
  }
  endTimeValid = (rule, value, callback) => {
    const { form } = this.props
    const startTime = form.getFieldValue('start_time')
    if (startTime === undefined) {
      return callback('请先输入开始时间')
    }
    if (moment(startTime).valueOf() > moment(value).valueOf()) {
      return callback('结束时间必须晚于开始时间')
    }
    callback()
  }
  cycleValid = (rule, value, callback) => {
    const { getFieldValue } = this.props.form
    if (value || value === 0) {
      if (parseInt(value) !== value) {
        // console.log(value, typeof value)
        return callback('请输入整数')
      }
      if (value < 1) {
        return callback('请输入大于 1 的整数')
      }
    }
    callback()
  }
  delayValid = (rule, value, callback) => {
    const { getFieldValue } = this.props.form
    if (value || value === 0) {
      if (parseInt(value) !== value) {
        // console.log(value, typeof value)
        return callback('请输入整数')
      }
      if (value < 0) {
        return callback('请输入大于等于 0 的整数')
      }
    }
    callback()
  }
  groupChange = () => {
    setTimeout(() => {
      // 异步可以使用getFieldValue拿到最新值
      const { validateFields, getFieldValue } = this.props.form
      const user = getFieldValue('user')
      const duty_group = getFieldValue('duty_group')
      const group = getFieldValue('group')
      this.setState({
        groupState: {
          user: !(user || duty_group || duty_group === 0 || group || group === 0),
          duty_group: !(user || duty_group || duty_group === 0 || group || group === 0),
          group: !(user || duty_group || duty_group === 0 || group || group === 0),
        },
      }, () => {
        validateFields(['user', 'duty_group', 'group'], { force: true }, (error, values) => {
          // console.log('error', error)
        })
      })
    })
  }
  render() {
    const { getFieldDecorator } = this.props.form
    const { mode, visible, groupState } = this.state

    return (
      <Modal
        title={mode === 'edit' ? '编辑报警策略' : '添加报警策略'}
        visible={visible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal" onValuesChange={this.groupChange}>
          <Form.Item label="报警时间段" style={{ marginBottom: 0 }}>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('start_time', {
                rules: [{ type: 'object', required: true, message: 'Please select time!' }],
              })(<TimePicker style={{ width: '100%' }} format="HH:mm" onChange={this.startTimeChange} />)}
            </Form.Item>
            <span style={{ display: 'inline-block', width: '20px', textAlign: 'center' }}>~</span>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('end_time', {
                rules: [
                  { type: 'object', required: true, message: 'Please select time!' },
                  // { validator: this.endTimeValid}
                ],
              })(<TimePicker style={{ width: '100%' }} format="HH:mm" />)}
            </Form.Item>
          </Form.Item>
          <Form.Item label="报警延迟">
            {getFieldDecorator('start', {
              rules: [
                { required: true, message: '请输入报警延迟' },
                { validator: this.delayValid },
              ],
            })(<InputNumber type="number" style={{ width: '100%' }} />)}
          </Form.Item>
          <Form.Item label="报警周期"
            wrapperCol={{
              xs: { span: 24 },
              sm: { span: 16 },
            }}
          >
            {getFieldDecorator('period', {
              rules: [
                { required: true, message: '请输入报警周期' },
                { validator: this.cycleValid },
              ],
            })(<InputNumber type="number" style={{ width: 'calc(100% - 20px)' }} />)}
            <span style={{ width: '20px', display: 'inline-block', textAlign: 'right' }}>分</span>
          </Form.Item>
          <Form.Item label="报警用户">
            {getFieldDecorator('user', {
              rules: [
                { required: groupState.user, message: '请输入报警用户' },
              ],
            })(<DInput onChange={this.groupChange} />)}
          </Form.Item>
          <Form.Item label="值班组">
            {getFieldDecorator('duty_group', {
              rules: [
                { required: groupState.duty_group, message: '请输入报警用户' },
              ],
            })(<DInput onChange={this.groupChange} />)}
          </Form.Item>
          <Form.Item label="报警用户组">
            {getFieldDecorator('group', {
              rules: [
                { required: groupState.group, message: '请输入报警用户组' },
              ],
            })(<DInput onChange={this.groupChange} />)}
          </Form.Item>
          <Form.Item label="Filter表达式">
            {getFieldDecorator('expression', {})(<Input />)}
          </Form.Item>
          <Form.Item label="报警方式">
            {getFieldDecorator('method', {
              initialValue: 'LANXIN',
              rules: [
                { required: true, message: '请输入报警用户组' },
              ],
            })(<Select onChange={this.selectMethod}>
              <Option value="LANXIN">LANXIN</Option>
              <Option value="CALL">CALL</Option>
              <Option value="SMS">SMS</Option>
              <Option value="HOOK">HOOK</Option>
            </Select>)}
          </Form.Item>
          {this.state.method === "HOOK" ?
            <Form.Item label="HOOK URL">
              {getFieldDecorator('hookurl', {
                // initialValue: hookurl,
                rules: [
                  { required: true, message: '请输入HOOK URL' },
                ],
              })(<Input />)}
            </Form.Item> : null
          }
        </Form>
      </Modal>
    )
  }
}
