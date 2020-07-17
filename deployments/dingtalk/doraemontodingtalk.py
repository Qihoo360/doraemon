import copy
import json
import os
import time
import hmac
import requests
import hashlib
import base64
import urllib.parse

from flask import Flask, jsonify, request

app = Flask(__name__)
app.debug = True


def senddingtalk(message, *args):
    timestamp = str(round(time.time() * 1000))
    secret_enc = secret.encode('utf-8')
    string_to_sign = '{}\n{}'.format(timestamp, secret)
    string_to_sign_enc = string_to_sign.encode('utf-8')
    hmac_code = hmac.new(secret_enc, string_to_sign_enc, digestmod=hashlib.sha256).digest()
    sign = urllib.parse.quote_plus(base64.b64encode(hmac_code))
    sign_url = url + "&timestamp="+timestamp+"&sign="+sign
    header = {
        "Content-Type": "application/json",
        "Charset": "UTF-8"
    }
    return requests.post(sign_url, json.dumps(message), headers=header).content




def alert_type(dictss):
    #  处理autorelease类型数据，并且格式化为dingtalk的格式
    msgautorelease = {
        "actionCard": {
            "title": "告警通知",
            "text": "var",
            "btnOrientation": "0",
        },
        "msgtype": "actionCard"
    }
    alert_name=dictss['alerts'][0]['summary']  #告警名称
    alert_active_time = dictss['time'] #告警发生时间
    alert_list=[] #存储格式化处理过的每一个alert
    if dictss['type'] == 'alert':
        alert_web_url=dictss['confirm_link'] #告警web url
        alert_str="""<font color=#FF3030 size=2 >告警通知 </font>  \n\n**告警时间**:"""+alert_active_time+"""\n\n**告警类别**:"""+alert_name+"""\n\n**AlertWeb**:"""+alert_web_url+"""\n\n**告警详情**:"""
        alert_status="告警通知"
        
    elif dictss['type']=='recover':
        alert_str="""<font color=#008000 size=2 >告警恢复 </font>  \n\n**恢复时间**:"""+alert_active_time+"""\n\n**恢复类别**:"""+alert_name+"""\n\n**恢复详情**:"""
        alert_status="告警恢复"
    else:
        pass



    for t in dictss['alerts']:     #删除无用标签,并格式化告警子条目
        #del_labeles需要清理的无用label,根据实际情况填写全局变量
        for del_labele in del_labeles:
            if del_labele in t['labels']:
                del t['labels'][del_labele]

        alert_id=str(t['id'])   #告警id
        alert_count_time=str(t['count']) #告警持续时间
        if len(t['hostname'])==0:
           alert_hostname="non_node"
        else:
           alert_hostname=t['hostname'] #告警主机，可为空

        alert_current_v=str(t['value']) #当前监控
        alert_label=str(t['labels'])
        if dictss['type'] == 'alert':
            alert_t_str="""\n\n————————————————\n\n 告警id:  """ +alert_id+ """    持续时间:""" +alert_count_time+"""m"""+"""\n\n告警主机: """+alert_hostname+"""\n\n当前监控值: """+alert_current_v+"""\n\nlabeles:\n\n"""+alert_label
        elif dictss['type']=='recover':
            alert_t_str="""\n\n————————————————\n\n 恢复告警id: """ +alert_id+"""\n\n恢复主机:"""+alert_hostname+"""\n\n当前监控值:  """+alert_current_v+"""\n\nlabeles:\n\n"""+alert_label
        else:
            pass
        alert_list.append(alert_t_str)
    
    for i in alert_list:  #所有告警条目汇总成一条
        alert_str+=i
    autorelease_msg = copy.deepcopy(msgautorelease)
    autorelease_msg["actionCard"]["text"] = alert_str
    autorelease_msg["actionCard"]["title"] = envname+alert_status
    return autorelease_msg

@app.route('/alert', methods=['post'])
def alert():
    if not request.data:  # 检测是否有数据
        return ('fail')
    global envname, url,secret,del_labeles
    envname = os.getenv('envname')
    url = os.getenv('dingtalk_url')
    secret = os.getenv('secret') 
    del_labeles=['heritage','chart','component','app','helm_sh_chart','app_kubernetes_io_managed_by','app_kubernetes_io_instance','kubernetes_node','app_kubernetes_io_name','kubernetes_name','job','release','kubernetes_namespace']
    student = request.data.decode('utf-8')
    # 获取到POST过来的数据，因为我这里传过来的数据需要转换一下编码。根据晶具体情况而定
    dictss = json.loads(student)
    print(dictss)
    print('\n')
    file = 'alert.json'
    with open(file, "a+") as f:
        f.write(json.dumps(dictss)+"\n"+"\n")
    alert_type_msg = alert_type(dictss)
    senddingtalk(alert_type_msg)

    return ('OK')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8100)
