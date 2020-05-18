import React, { Component } from 'react'
import { Modal, message, Form, Input, Icon } from 'antd'
import { formItemLayout } from '@configs/const'

@Form.create({})
export default class ChangeUserPassword extends Component {
    state = {
        // labels: [0],
        id: 0,
        visiable: false,
    }
    componentDidMount() {
        this.props.onRef(this)
    }
    componentWillReceiveProps() {

    }
    updateValue(value) {
        const { form } = this.props
        form.resetFields();
        this.setState({
            id: value ? value.id : 0,
            visiable: true,
        })
        value && form.setFieldsValue(value)
    }

    handleOk = (e) => {
        this.props.form.validateFields(async (err, values) => {
            // console.log(values)
            if (!err) {
                const { id } = this.state;
                const resultSuccess = await this.props.onSubmit({ id, ...values })
                if (resultSuccess) {
                    if (id) {
                        message.success('修改成功')
                    } else {
                        message.success('密码修改成功')
                    }
                    this.setState({
                        visiable: false,
                    })
                }
            }
        })
    }
    handleCancel = (e) => {
        const { form } = this.props
        form.resetFields()
        this.setState({
            visiable: false,
        })
    }
    // addLabel = (e) => {
    //   const { labels } = this.state
    //   const last = labels.length - 1
    //   const nextKeys = labels.concat(last + 1)
    //   this.setState({
    //     labels: nextKeys,
    //   })
    // }
    // removeLabel = (k) => {
    //   const { labels } = this.state
    //   if (labels.length === 1) { return }
    //   this.setState({
    //     labels: labels.filter(key => key !== k),
    //   })
    // }
    render() {
        const { getFieldDecorator } = this.props.form
        const { id, visiable } = this.state
        return (
            <Modal
                title={id ? '编辑报警组' : '修改登录密码'}
                visible={visiable}
                onOk={this.handleOk}
                onCancel={this.handleCancel}
                maskClosable={false}
            >
                <Form {...formItemLayout} layout="horizontal">
                    <Form.Item label="原密码">
                        {getFieldDecorator('oldpassword', {
                            rules: [
                                { required: true, message: '请输入原密码' },
                            ],
                        })(<Input prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            type="password" />)}
                    </Form.Item>
                    <Form.Item label="新密码">
                        {getFieldDecorator('newpassword', {
                            rules: [
                                { required: true, message: '请输入新密码' },
                            ],
                        })(<Input prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            type="password" />)}
                    </Form.Item>
                    <Form.Item label="确认新密码">
                        {getFieldDecorator('confirmnewpassword', {
                            rules: [
                                { required: true, message: '请输入确认新密码' },
                            ],
                        })(<Input prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            type="password" />)}
                    </Form.Item>
                </Form>
            </Modal>
        )
    }
}
