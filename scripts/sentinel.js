const fs = require('fs');
const path = require('path');

function checkStandards(basePath) {
    const standardsPath = path.join(basePath, 'docs', 'standards');
    const requiredFiles = [
        'architecture.md', 'conventions.md', 'workflow.md', 
        'status-protocol.md', 'contribution.md', 'core-contribution.md'
    ];
    const errors = [];
    if (!fs.existsSync(standardsPath)) return ["Docs standards directory missing"];
    
    requiredFiles.forEach(f => {
        if (!fs.existsSync(path.join(standardsPath, f))) {
            errors.push(`Missing standard file: ${f}`);
        }
    });

    const repoPlanPath = path.join(basePath, 'docs', 'repository-plan', 'README.md');
    if (!fs.existsSync(repoPlanPath)) {
        errors.push("Missing repository plan blueprint: docs/repository-plan/README.md");
    }

    return errors;
}

function auditAdvancedREADME(filePath, errors) {
    const content = fs.readFileSync(filePath, 'utf8');
    const isExplicitNil = content.includes('Unit ini tidak membutuhkan Lab Praktis/Visualisasi');
    
    const requirements = [
        { name: "Source/Clause Anchor", pattern: /Source:|Clause/i, optionalIfNil: false },
        { name: "Logic-Pure Definition", pattern: /Logic-Pure/i, optionalIfNil: false },
        { name: "Analogy-Model", pattern: /Analogy-Model/i, optionalIfNil: false },
        { name: "Internal Mechanics/Algorithm", pattern: /Mekanisme Internal|Algoritma|Steps/i, optionalIfNil: true },
        { name: "State/Architecture Mapping", pattern: /State & Architecture Mapping|Internal Slots|Slot/i, optionalIfNil: true },
        { name: "Deep-Visual (Mermaid)", pattern: /```mermaid/i, optionalIfNil: true },
        { name: "Experimental Lab Link", pattern: /Lab Praktis|examples\/|Experimental Lab/i, optionalIfNil: true },
        { name: "Cross-Rack Linking", pattern: /Hubungan Sistem|Cross-Rack|RAK-/i, optionalIfNil: true }
    ];

    requirements.forEach(req => {
        if (!req.pattern.test(content)) {
            if (!(isExplicitNil && req.optionalIfNil)) {
                errors.push(`[ADVANCED-ERROR] Missing ${req.name} in ${filePath}`);
            }
        }
    });
}

function auditStructure(basePath) {
    const errors = [];
    function walk(dir) {
        const files = fs.readdirSync(dir);
        files.forEach(file => {
            const fullPath = path.join(dir, file);
            if (fs.statSync(fullPath).isDirectory()) {
                if (file.startsWith('.') || file === 'node_modules' || file === 'scripts' || file === 'assets' || file === 'examples') return;
                
                if (file.startsWith('RAK-') || file.startsWith('SR-') || file.startsWith('BK-') || file.startsWith('CH-') || file.startsWith('SEC-')) {
                    const readmePath = path.join(fullPath, 'README.md');
                    if (!fs.existsSync(readmePath)) {
                        errors.push(`Missing README.md in ${fullPath}`);
                    } else {
                        // Audit logic for CH- and SEC-
                        if (file.startsWith('CH-') || file.startsWith('SEC-')) {
                            const rakMatch = fullPath.match(/RAK-(\d+)/);
                            const isAdvanced = rakMatch && (parseInt(rakMatch[1]) >= 2);
                            
                            if (isAdvanced) {
                                auditAdvancedREADME(readmePath, errors);
                            }
                        }
                    }
                }
                walk(fullPath);
            }
        });
    }
    walk(basePath);
    return errors;
}

function main() {
    const basePath = path.dirname(__dirname);
    console.log(`--- Sentinel Audit (Unified 6-Rack Edition) ---`);
    console.log(`Auditing: ${basePath}\n`);
    
    const errors = [...checkStandards(basePath), ...auditStructure(basePath)];
    
    if (errors.length === 0) {
        console.log("[PASS] Everything is perfectly standardized to the Unified 6-Rack Gold Standard! 🏆");
    } else {
        console.log(`[FAIL] Found ${errors.length} inconsistencies:`);
        errors.forEach(err => console.log(` - ${err}`));
        process.exit(1);
    }
}

main();
