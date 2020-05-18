import React, { Component } from 'react'
import { Tag } from 'antd'
import { keyCode } from '@configs/const'
// import { MyContext } from '@content/group/context'

import './input.less'

// const MyContext = React.createContext('defaultValue')
/**
@props separator 分隔符
*/
const defaultSeparator = ','

class DInput extends Component {
  // static contextType = MyContext
  constructor(props, context) {
    super(props, context)
    console.log('context', context)
  }
  state = {
    showPlaceholder: true,
  }
  componentDidMount() {
    // console.log(MyContext)
    console.log(this.context)
  }
  componentDidUpdate() {
    this.input.style.width = '2px'
  }
  updateChange({ newValue, add }) {
    let { separator, value } = this.props
    separator = separator || defaultSeparator
    value = value || ''
    this.props.onChange(newValue !== undefined ? newValue : `${value ? value + separator : ''}${add}`)
  }
  containerClick = () => {
    this.input.focus()
    this.setState({
      showPlaceholder: false,
    })
  }
  refInput = (c) => {
    this.input = c
  }
  closeTag = (e, tag) => {
    e.preventDefault();
    let { separator } = this.props
    const { value } = this.props
    separator = separator || defaultSeparator
    const newValue = value.split(separator).filter(item => item !== tag).join(separator)
    this.updateChange({ newValue })
  }
  inputKeyDown = (e) => {
    this.input.style.width = `${this.input.scrollWidth + 10}px`
    const { value } = e.target
    if (e.keyCode === keyCode.Enter) {
      this.enterEvent(value)
      this.resetInput(e.target)
      this.input.focus()
    }
  }
  enterEvent(value) {
    value && this.updateChange({ add: value })
  }
  inputBlur = (e) => {
    const { value } = e.target
    this.enterEvent(value)
    this.resetInput(e.target)
    this.setState({
      showPlaceholder: true,
    })
  }
  resetInput(target) {
    target.value = ''
  }
  render() {
    // const { list } = this.state
    let { value, separator } = this.props
    // 这里重置为 undefined 不能改变input的值
    value = value || ''
    separator = separator || ','
    const list = value.split(separator).filter(item => item !== '')
    return (
      <div>
        {/* <MyContext.Provider value="dark">
          context: {this.context}
        </MyContext.Provider>
        context: {this.context} / {MyContext._currentValue} */}
        <div className="d-input-container" onClick={this.containerClick}>
          {(list.length === 0 && this.state.showPlaceholder) && (<div className="d-input-placeholder">按回车键输入</div>)}
          <ul>
            {list.map((item, index) => (<li><Tag key={item} style={{ marginRight: 5, marginBottom: 5 }} closable onClose={e => this.closeTag(e, item)}>
              {item}
            </Tag></li>))}
            <li>
              <input ref={this.refInput} type="text" onKeyDown={this.inputKeyDown} onBlur={e => this.inputBlur(e)} />
            </li>
          </ul>
        </div>
      </div>
    )
  }
}
// DInput.contextType = MyContext
export { DInput }
