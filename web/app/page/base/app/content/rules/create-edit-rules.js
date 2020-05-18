import React, { Component } from 'react'
import { Modal, message, Form, Input, Select } from 'antd'
import { formItemLayout } from '@configs/const'
import { promethus, strategy } from '@actions/common'
import { connect } from 'react-redux'

const { Option } = Select;

@connect((state, props) => ({
  promethusState: state.promethus.loaded,
  promethusData: state.promethus.data,
  strategyState: state.strategy.loaded,
  strategyData: state.strategy.data,
}), dispatch => ({
  promethusAction: () => dispatch(promethus()),
  strategyAction: () => dispatch(strategy()),
}))
@Form.create({})
export default class CreateEditRules extends Component {
  constructor(props) {
    super(props)
    const { promethusState, strategyState, promethusAction, strategyAction } = props
    if (!promethusState) { promethusAction() }
    if (!strategyState) { strategyAction() }
  }
  state = {
    id: 0,
  }
  componentDidMount() {
    this.props.onRef(this)
  }
  updateValue(value) {
    const { form } = this.props
    form.resetFields();
    this.setState({
      id: value ? value.id : 0,
    })
    form.setFieldsValue(value)
  }

  handleOk = (e) => {
    this.props.form.validateFields(async (err, values) => {
      if (!err) {
        const { id } = this.state;
        const resultSuccess = await this.props.onSubmit({ id, ...values })
        if (resultSuccess) {
          if (id) {
            message.success('修改成功')
          } else {
            message.success('添加成功')
          }
          this.setState({
            id: 0,
          })
        }
      }
    })
  }
  handleCancel = (e) => {
    this.props.onClose()
    const { form } = this.props
    form.resetFields()
    this.setState({
      id: 0,
    })
  }
  forChange = (e) => {
    const { value } = e.target
    e.target.value = `${parseFloat(value) || ''}s`
  }
  render() {
    const { visiable, promethusData, strategyData } = this.props
    const { getFieldDecorator } = this.props.form
    const { id } = this.state
    const selectBefore = getFieldDecorator('op', {
      initialValue: '>',
    })(<Select style={{ width: 70 }}>
      <Option value="==">==</Option>
      <Option value="!=">!=</Option>
      <Option value=">">&gt;</Option>
      <Option value="<">&lt;</Option>
      <Option value=">=">&gt;=</Option>
      <Option value="<=">&lt;=</Option>
    </Select>);
    return (
      <Modal
        title={id ? '编辑报警规则管理' : '添加报警规则管理'}
        visible={visiable}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="监控指标">
            {getFieldDecorator('expr', {
              rules: [
                { required: true, message: '请输入监控指标' },
              ],
            })(<Input />)}
          </Form.Item>
          <Form.Item label="报警阈值">
            {getFieldDecorator('value', {
              rules: [
                { required: true, message: '请输入报警阈值' },
              ],
            })(<Input addonBefore={selectBefore} />)}
          </Form.Item>
          <Form.Item label="持续时间">
            {getFieldDecorator('for', {
              initialValue: '0s',
              rules: [
                { required: true, message: '请输入持续时间' },
              ],
            })(<Input onChange={this.forChange} />)}
          </Form.Item>
          <Form.Item label="标题">
            {getFieldDecorator('summary', {
              rules: [
                { required: true, message: '请输入标题' },
              ],
            })(<Input />)}
          </Form.Item>
          <Form.Item label="描述">
            {getFieldDecorator('description', {
              rules: [
              ],
            })(<Input />)}
          </Form.Item>
          <Form.Item label="数据源">
            {getFieldDecorator('prom_id', {
              rules: [
                { required: true, message: '请输入数据源' },
              ],
            })(<Select style={{ width: '100%' }}>
              {
                promethusData && promethusData.map(item => (
                  <Option value={item.id}>{item.name}</Option>
                ))
              }
            </Select>)}
          </Form.Item>
          <Form.Item label="策略">
            {getFieldDecorator('plan_id', {
              rules: [
                { required: true, message: '请输入策略' },
              ],
            })(<Select style={{ width: '100%' }}>
              {
                strategyData && strategyData.map(item => (
                  <Option value={item.id}>{item.description}</Option>
                ))
              }
            </Select>)}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}
