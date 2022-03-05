import pyperclip
import json
import requests
import keyboard

book = ""



def translator(data):
    """
    input : str 需要翻译的字符串
    output：translation 翻译后的字符串
    """
    # API
    url = 'http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule&smartresult=ugc&sessionFrom=null'
    # 传输的参数， i为要翻译的内容
    key = {
        'type': "AUTO",
        'i': data,
        "doctype": "json",
        "version": "2.1",
        "keyfrom": "fanyi.web",
        "ue": "UTF-8",
        "action": "FY_BY_CLICKBUTTON",
        "typoResult": "true"
    }
    # key 这个字典为发送给有道词典服务器的内容
    response = requests.post(url, data=key)
    # 判断服务器是否相应成功
    if response.status_code == 200:
        # 通过 json.loads 把返回的结果加载成 json 格式
        result = json.loads(response.text)
#         #print ("输入的词为：%s" % result['translateResult'][0][0]['src'])
#         #print ("翻译结果为：%s" % result['translateResult'][0][0]['tgt'])
        translation = result['translateResult'][0][0]['tgt']
        return translation
    else:
        print("有道词典调用失败")
        # 相应失败就返回空
        return None


#这里应该可以优化一下，当剪贴板内容被更新之后再去创建一个新的waitforNewPast,而不是每一个循环都创建一次
def trans():
    global book
    data=pyperclip.waitForNewPaste()
    translated = translator(data)
    tmp = data+" : "+translated
    print("Elina> "+tmp)
    book = book+"\n"+tmp
def export():
    global book
    filename = 'book.txt'
    with open(filename, 'w') as file_object:
        file_object.write(book)
    print("Elina> 已导出" )


keyboard.add_hotkey('ctrl+c', trans)
# 按f1输出aaa
keyboard.add_hotkey('ctrl+o', export)
# 按ctrl+alt输出b
keyboard.wait('ctrl+e') #end



