import * as monaco from "monaco-editor";

export function createLaunchArgumentsEditor(domElement, launchArguments) {
    const editor = monaco.editor.create(domElement, {
        value: launchArguments.join('\n'),
        language: 'ini',
        lineNumbers: 'off',
        theme: 'vs-dark',
        minimap: {
            enabled: false // 是否启用预览图
        },
        automaticLayout: true,
        scrollBeyondLastLine: false,
    });
    editor.addAction({
        id: "id1",
        label: "填充示例",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 1.5,
        run: function (ed) {
            editor.setValue("-AudioMixer\n-RenderOffScreen\n-ForceRes\n-ResX=1920\n-ResY=1080")
        },
    })
    editor.addAction({
        id: "id2",
        label: "还原更改",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 2,
        run: function (ed) {
            editor.setValue(launchArguments.join('\n'))
        },
    })
    return editor
}

export function createMetadataEditor(domElement, metadata) {
    const editor = monaco.editor.create(domElement, {
        value: metadata,
        language: 'yaml',
        lineNumbers: 'off',
        theme: 'vs-dark',
        minimap: {
            enabled: false // 是否启用预览图
        },
        automaticLayout: true,
        scrollBeyondLastLine: false
    })
    editor.getModel().updateOptions({tabSize: 2})
    editor.addAction({
        id: "id1",
        label: "填充示例",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 1.5,
        run: function (ed) {
            editor.setValue("labels: #以下是key: value格式\n  key1: value1\n  key2: value2")
        },
    })
    editor.addAction({
        id: "id2",
        label: "还原更改",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 2,
        run: function (ed) {
            editor.setValue(metadata)
        },
    })
    return editor
}

export function createPaksConfigEditor(domElement, paksConfig) {
    const editor = monaco.editor.create(domElement, {
        value: paksConfig,
        language: 'yaml',
        lineNumbers: 'off',
        theme: 'vs-dark',
        minimap: {
            enabled: false // 是否启用预览图
        },
        automaticLayout: true,
        scrollBeyondLastLine: false
    })
    editor.getModel().updateOptions({tabSize: 2})
    editor.addAction({
        id: "id1",
        label: "填充示例",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 1.5,
        run: function (ed) {
            editor.setValue("paks:\n  - name: 宜宾换流站    #列表里显示名称\n    value: yibin       #pak目录名称 \n" +
                "  - name: 雁门关换流站\n    value: yanmenguan\n  - name: 中都换流站\n    value: zhongdu")
        },
    })
    editor.addAction({
        id: "id2",
        label: "还原更改",
        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: "navigation",
        contextMenuOrder: 2,
        run: function (ed) {
            editor.setValue(paksConfig)
        },
    })
    return editor
}